package main;

import "fmt";

func binarySearch(arr []int, n int) bool {
  found := false
  start := 0
  end := len(arr)

  for !found && start <= end {
    mid := (start + end) / 2

    if arr[mid] == n {
      found = true
    } else if arr[mid] < n {
      start = mid + 1
    } else {
      end = mid - 1
    }
  }

  return found
}

func sequentialSearch(arr []int, n int) bool {
  for i := 0; i < len(arr); i++ {
    if(arr[i] == n) {
      return true;
    }
  }

  return false;
}

func main() {
  var array = []int{1,3,5}
  n := 1
  fmt.Println(binarySearch(array, n))
  fmt.Println(sequentialSearch(array, n))
}
