package main

import (
	"errors"
	"fmt"
	"github.com/howeyc/fsnotify"
	"os"
	"path/filepath"
	"time"
)

// watchAllDirs will watch a directory and then walk through all subdirs
// of that directory and watch them too
func watchAllDirs(watcher *fsnotify.Watcher, root string) (err error) {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			watcher.Watch(path)
		}
		return nil
	}
	return filepath.Walk(root, walkFn)
}

// Watch watches a path for creation, deletion, modification and renaming and
// will execute command.Execute() when such an event occurs
// if path is a folder, it will watch the folder itself and all its content
// including subfolders
func Watch(path string, command Command) error {
	done := make(chan bool)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.New(fmt.Sprintf("Was not able to watch %s", path))
	}

	go func() {
		eventstream := watcher.Event
		start_time := time.Now()
		for event := range eventstream {
			// 800 milliseconds is an approximate minimum we use because
			// when a file is edited, is will send multiple events with
			// different types but we only want to get notified one
			//
			// id very much like a cleaner more sensible solution :)
			if time.Since(start_time) > time.Millisecond*800 {
				// if a subfolder is added we need to watch it and its content also
				if event.IsCreate() {
					if finfo, err := os.Stat(event.Name); err == nil && finfo.IsDir() {
						watcher.Watch(event.Name)
					}
				}
				command.Execute()
				start_time = time.Now()
			}
		}
		done <- true
	}()

	if finfo, err := os.Stat(path); err == nil && finfo.IsDir() {
		err = watchAllDirs(watcher, path)
	} else {
		err = watcher.Watch(path)
	}

	if err != nil {
		return errors.New(fmt.Sprintf("Was not able to watch %s", path))
	}

	<-done

	watcher.Close()

	return nil
}
