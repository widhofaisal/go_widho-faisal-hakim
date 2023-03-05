# Section 10 - Recursive - Number Theory - Sorting - Searching

## Video – Recursive – Number Theory – Sorting – Searchings (15:52)

<br/>

### **Recursive**
```golang
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	fmt.Println(factorial(5))
}
```

<br/>

### **Number Theory**
- bilangan prima
  <p>ex : 5! = 5x4x3x2x1 = 120</p>
- FPB
- KPK
- factorial
- faktor prima
- linier search O(n)
- built in search
- sorting
  - selection sort O(n^2)
  - counting sort O(n+k)
- built in sorting