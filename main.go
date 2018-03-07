package main

func main() {
	bc := NewBlockchain("0x000000000000000000000000000")
	defer bc.DB.Close()

	cli := CLI{bc}
	cli.Run()
}
