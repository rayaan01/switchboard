package db

import (
	"time"

	"github.com/google/uuid"
)

func measure_tps_set(engine StorageEngine) (int, error) {
	var count int
	duration := time.Second
	done := make(chan bool)

	go func() {
		time.Sleep(duration)
		done <- true
	}()

	for {
		select {
		case <-done:
			return count, nil
		default:
			key := uuid.NewString()
			value := uuid.NewString()
			_, err := engine.set(key, value)
			if err != nil {
				return -1, err
			}
			count++
		}
	}
}

func measure_tps_get(engine StorageEngine, keys []string) (int, error) {
	var count int
	duration := time.Second
	done := make(chan bool)

	go func() {
		time.Sleep(duration)
		done <- true
	}()

	for _, key := range keys {
		select {
		case <-done:
			return count, nil
		default:
			_, err := engine.get(key)
			if err != nil {
				return -1, err
			}
			count++
		}
	}

	return -1, nil
}

func measure_tps_del(engine StorageEngine, keys []string) (int, error) {
	var count int
	duration := time.Second
	done := make(chan bool)

	go func() {
		time.Sleep(duration)
		done <- true
	}()

	for _, key := range keys {
		select {
		case <-done:
			return count, nil
		default:
			_, err := engine.del(key)
			if err != nil {
				return -1, err
			}
			count++
		}
	}

	return -1, nil
}
