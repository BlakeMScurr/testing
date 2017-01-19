### Setup

If you have golang setup on your computer, or you are happy to do so ([use this](http://tecadmin.net/install-go-on-ubuntu/) if you are using ubuntu) please fork and clone this repo and submit a pull request to solve the problems below.

```git clone https://github.com/BlakeMScurr/testing```

If you would rather get straight to the problems, use [these](https://play.golang.org/p/bMD3B1JPAw) [playground](https://play.golang.org/p/6rSAir4oim) links.

### Task

hello_world.go:

When you run `$ go run hello_world.go` you should have the output "hello world 1 2 3"

recursion.go:

When you run `$ go run recursion.go n` you should get the nth fibonnaci number.
For example, `$ go run recursion.go 10` should give 55
Also, `go run recursion.go 0` should give 0.
But you are not expected to deal with negative numbers.

Once that's working, try altering the algorithm so it can quickly find numbers above 50.
(The new algorithm doesn't have to be recursive).
