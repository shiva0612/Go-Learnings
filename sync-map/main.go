package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	// ===========================================
	// 1. Store - Set a key-value pair
	// ===========================================
	m.Store("name", "Alice")
	m.Store("age", 30)
	m.Store(123, "numeric key")

	fmt.Println("--- Store ---")
	fmt.Println("Stored: name=Alice, age=30")

	// ===========================================
	// 2. Load - Get a value by key
	// ===========================================
	fmt.Println("\n--- Load ---")
	if value, ok := m.Load("name"); ok {
		fmt.Printf("name: %v\n", value)
	}

	if _, ok := m.Load("missing"); !ok {
		fmt.Println("'missing' key not found")
	}

	// ===========================================
	// 3. LoadOrStore - Get existing or store new
	// ===========================================
	fmt.Println("\n--- LoadOrStore ---")

	// First call - stores the value
	actual, loaded := m.LoadOrStore("city", "New York")
	fmt.Printf("city: %v, was already present: %v\n", actual, loaded)

	// Second call - returns existing value
	actual, loaded = m.LoadOrStore("city", "Los Angeles")
	fmt.Printf("city: %v, was already present: %v\n", actual, loaded)

	// ===========================================
	// 4. LoadAndDelete - Get and remove atomically
	// ===========================================
	fmt.Println("\n--- LoadAndDelete ---")

	m.Store("temp", "temporary value")
	if value, loaded := m.LoadAndDelete("temp"); loaded {
		fmt.Printf("Deleted temp: %v\n", value)
	}

	// Verify it's deleted
	if _, ok := m.Load("temp"); !ok {
		fmt.Println("temp key no longer exists")
	}

	// ===========================================
	// 5. Delete - Remove a key-value pair
	// ===========================================
	fmt.Println("\n--- Delete ---")

	m.Store("remove_me", "will be deleted")
	m.Delete("remove_me")
	fmt.Println("Deleted 'remove_me' key")

	// ===========================================
	// 6. Range - Iterate over all entries
	// ===========================================
	fmt.Println("\n--- Range ---")

	m.Store("user1", "Bob")
	m.Store("user2", "Carol")
	m.Store("user3", "Dave")

	fmt.Println("All entries:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // continue iteration
	})

	// Range with early termination
	fmt.Println("\nFirst 2 entries only:")
	count := 0
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		count++
		return count < 2 // stop after 2 items
	})

	// ===========================================
	// 7. CompareAndSwap - Conditional update (Go 1.20+)
	// ===========================================
	fmt.Println("\n--- CompareAndSwap ---")

	m.Store("counter", 10)

	// Try to swap - succeeds because old value matches
	swapped := m.CompareAndSwap("counter", 10, 20)
	fmt.Printf("Swap 10->20: %v\n", swapped)

	// Try to swap - fails because old value doesn't match
	swapped = m.CompareAndSwap("counter", 10, 30)
	fmt.Printf("Swap 10->30: %v (current value is 20)\n", swapped)

	if value, _ := m.Load("counter"); value != nil {
		fmt.Printf("Final counter value: %v\n", value)
	}

	// ===========================================
	// 8. CompareAndDelete - Conditional delete (Go 1.20+)
	// ===========================================
	fmt.Println("\n--- CompareAndDelete ---")

	m.Store("status", "active")

	// Try to delete - fails because value doesn't match
	deleted := m.CompareAndDelete("status", "inactive")
	fmt.Printf("Delete if 'inactive': %v\n", deleted)

	// Try to delete - succeeds because value matches
	deleted = m.CompareAndDelete("status", "active")
	fmt.Printf("Delete if 'active': %v\n", deleted)

	// ===========================================
	// 9. Swap - Replace value and return old (Go 1.20+)
	// ===========================================
	fmt.Println("\n--- Swap ---")

	m.Store("version", "1.0")

	previous, loaded := m.Swap("version", "2.0")
	fmt.Printf("Old version: %v, existed: %v\n", previous, loaded)

	// Swap on non-existent key
	previous, loaded = m.Swap("new_key", "new_value")
	fmt.Printf("New key swap - old value: %v, existed: %v\n", previous, loaded)

	// ===========================================
	// 10. Concurrent Access Example
	// ===========================================
	fmt.Println("\n--- Concurrent Access ---")
	demonstrateConcurrency()

	// ===========================================
	// 11. Clear (Go 1.23+) - Remove all entries
	// ===========================================
	fmt.Println("\n--- Clear ---")
	m.Clear()
	fmt.Println("Map cleared")

	// Verify it's empty
	isEmpty := true
	m.Range(func(key, value interface{}) bool {
		isEmpty = false
		return false
	})
	fmt.Printf("Map is empty: %v\n", isEmpty)
}

func demonstrateConcurrency() {
	var m sync.Map
	var wg sync.WaitGroup

	// Multiple goroutines writing concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("goroutine_%d_item_%d", id, j)
				m.Store(key, id*10+j)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	// Multiple goroutines reading concurrently
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			m.Range(func(key, value interface{}) bool {
				// Just iterate
				return true
			})
		}(i)
	}

	wg.Wait()

	// Count final entries
	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("Total entries after concurrent operations: %d\n", count)
}
