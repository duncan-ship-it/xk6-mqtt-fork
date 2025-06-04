package mqtt

func getSize(value interface{}) int {
	switch v := value.(type) {
	case string:
		return len(v)
	case []byte:
		return len(v)
	case []string:
		return len(v)
	case []int:
		return len(v)
	case []int32:
		return len(v)
	case []int64:
		return len(v)
	case []uint:
		return len(v)
	case []uint16:
		return len(v)
	case []uint64:
		return len(v)
	case []float32:
		return len(v)
	case []float64:
		return len(v)
	case map[string]string:
		return len(v)
	case map[string]interface{}:
		return len(v)
	case map[int]string:
		return len(v)
	case map[int]int:
		return len(v)
	case chan int:
		return len(v)
	case chan struct{}:
		return len(v)
	case [4]int:
		return len(v)
	case [8]byte:
		return len(v)
	case *string:
		return len(*v)
	case *[]byte:
		return len(*v)
	case *[4]int:
		return len(*v)
	default:
		return 0
	}
}
