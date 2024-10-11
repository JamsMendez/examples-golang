package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	runBackup()
}

type BackupSystem struct {
	inProgress bool
	mu         sync.RWMutex
	rwmu       sync.RWMutex
}

func NewBackupSystem() *BackupSystem {
	return &BackupSystem{}
}

func (b *BackupSystem) Backup(ID int) {
	b.mu.Lock()

	if b.inProgress {
		fmt.Println("backup esta en progreso. Operacion rechazada ", ID)
		b.mu.Unlock()
		return
	}

	b.mu.Lock()
	b.inProgress = true
	b.mu.Unlock()
	fmt.Println("iniciando backup ... ", ID)

	<-time.After(3 * time.Second)

	b.mu.Lock()
	b.inProgress = false
	b.mu.Unlock()

	fmt.Println("backup completo ", ID)
}

func (b *BackupSystem) RWBackup(ID int) {
	b.rwmu.RLock()

	if b.inProgress {
		fmt.Println("backup esta en progreso. Operacion rechazada ", ID)
		b.rwmu.RUnlock()
		return
	}

	b.rwmu.RUnlock()

	b.rwmu.Lock()
	b.inProgress = true
	b.rwmu.Unlock()
	fmt.Println("iniciando backup ... ", ID)

	<-time.After(3 * time.Second)

	b.rwmu.Lock()
	b.inProgress = false
	b.rwmu.Unlock()

	fmt.Println("backup completo ", ID)
}

func (b *BackupSystem) UnsafeBackup(ID int) {
	if b.inProgress {
		fmt.Println("backup esta en progreso. Operacion rechazada ", ID)
		return
	}

	b.inProgress = true
	fmt.Println("iniciando backup ... ", ID)

	<-time.After(3 * time.Second)

	b.inProgress = false

	fmt.Println("backup completo ", ID)
}

func runBackup() {
	backup := NewBackupSystem()

	var wg sync.WaitGroup

	for i := 0; i < 10_000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			ID := index + 1
			backup.UnsafeBackup(ID)
		}(i)
	}

	wg.Wait()
}
