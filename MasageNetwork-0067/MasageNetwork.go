package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	masks := make([]uint, n)
	for i := 0; i < n; i++ {
		var mask string
		fmt.Scan(&mask)
		masks[i] = ipToInt(mask)
	}
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		var ip1, ip2 string
		fmt.Scan(&ip1, &ip2)
		ip1Int := ipToInt(ip1)
		ip2Int := ipToInt(ip2)
		count := 0
		for _, mask := range masks {
			if ip1Int&mask == ip2Int&mask {
				count++
			}
		}
		fmt.Println(count)
	}

}

func ipToInt(s string) uint {
	parts := strings.Split(s, ".")
	var result uint
	for i := 0; i < len(parts); i++ {
		curr, _ := strconv.Atoi(parts[i])
		result = (result << 8) | uint(curr)
	}
	return result
}

/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ipToInt converts an IP address (e.g., "192.168.0.1") to a 32-bit integer
func ipToInt(s string) uint32 {
	parts := strings.Split(s, ".")
	var result uint32
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		result = (result << 8) | uint32(num)
	}
	return result
}

func main() {
	// Use bufio for efficient I/O
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Read number of subnets
	var n int
	fmt.Fscan(in, &n)

	// Read and convert subnet masks to 32-bit integers
	masks := make([]uint32, n)
	for i := 0; i < n; i++ {
		var mask string
		fmt.Fscan(in, &mask)
		masks[i] = ipToInt(mask)
	}

	// Read number of IP pairs
	var m int
	fmt.Fscan(in, &m)

	// Process each IP pair
	for i := 0; i < m; i++ {
		var ip1, ip2 string
		fmt.Fscan(in, &ip1, &ip2)
		ip1Int := ipToInt(ip1)
		ip2Int := ipToInt(ip2)

		// Count subnets where IP1 and IP2 are in the same subnet
		count := 0
		for _, mask := range masks {
			if ip1Int&mask == ip2Int&mask {
				count++
			}
		}
		fmt.Fprintln(out, count)
	}
}*/
