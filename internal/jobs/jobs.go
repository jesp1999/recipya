package jobs

import (
	"github.com/go-co-op/gocron"
	"github.com/reaper47/recipya/internal/services"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"time"
)

// ScheduleCronJobs schedules cron jobs for the web app. It starts the following jobs:
//
// - Clean Images: Removes unreferenced images from the data/img folder to save space.
func ScheduleCronJobs(repo services.RepositoryService, imagesDir string) {
	scheduler := gocron.NewScheduler(time.UTC)
	_, _ = scheduler.Every(1).MonthLastDay().Do(func() {
		rmFunc := func(file string) error {
			return os.Remove(filepath.Join(imagesDir, file))
		}
		numFiles, numBytes := cleanImages(os.DirFS(imagesDir), repo.Images(), rmFunc)

		var s string
		if numBytes > 0 {
			s = "(" + strconv.FormatFloat(float64(numBytes)/(1<<20), 'f', 2, 64) + " MB)"
		}
		log.Printf("CleanImages: Removed %d unreferenced images %s", numFiles, s)
	})
	scheduler.StartAsync()
}

func cleanImages(dir fs.FS, usedImages []string, rmFileFunc func(path string) error) (numFilesDeleted, numBytesDeleted int64) {
	sort.Strings(usedImages)

	_ = fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		if path == "." {
			return nil
		}

		_, found := slices.BinarySearch(usedImages, d.Name())
		if !found {
			info, err := d.Info()
			if err != nil {
				log.Printf("clean images dir walk error: %s", err)
				return err
			}

			err = rmFileFunc(path)
			if err != nil {
				log.Printf("clean images walk '%s': %s", path, err)
				return err
			}

			numFilesDeleted++
			numBytesDeleted += info.Size()
		}
		return nil
	})
	return
}
