package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"strconv"
)

var (
	val string
	ttl int
	dv float64

	result int
	count int

	fresult float64
)
func main()  {
	myCal:=ui.Main(func() {

		input1:=ui.NewEntry()

		equl:=ui.NewButton("=")

		empty1:=ui.NewLabel("")
		empty2:=ui.NewLabel("")
		aempty:=ui.NewLabel("")

		empty6:=ui.NewLabel("  ")

		box1:=ui.NewVerticalBox()
		box1.Append(input1,false)
		box1.Append(empty1,false)

		box1.Append(empty2,false)

		box1.Append(aempty,false)

		plus:=ui.NewButton("+")
		minus:=ui.NewButton("-")
		mul:=ui.NewButton("*")
		divide:=ui.NewButton("/")
		ts:=ui.NewButton("Clear")

		box2:=ui.NewHorizontalBox()
		box2.Append(plus,false)

		box2.Append(minus,false)

		box2.Append(mul,false)

		box2.Append(divide,false)
		box2.Append(empty6,false)
		box2.Append(ts,false)

		seven:=ui.NewButton("7")
		eight:=ui.NewButton("8")
		nine:=ui.NewButton("9")

		four:=ui.NewButton("4")
		five:=ui.NewButton("5")
		six:=ui.NewButton("6")
		//
		one:=ui.NewButton("1")
		two:=ui.NewButton("2")
		three:=ui.NewButton("3")

		zero:=ui.NewButton("0")

		box3:=ui.NewHorizontalBox()
		box3.Append(seven,false)
		box3.Append(eight,false)
		box3.Append(nine,false)
		box3.Append(equl,false)

		box4:=ui.NewHorizontalBox()
		box4.Append(four,false)
		box4.Append(five,false)
		box4.Append(six,false)

		box5:=ui.NewHorizontalBox()
		box5.Append(one,false)
		box5.Append(two,false)
		box5.Append(three,false)

		box6:=ui.NewHorizontalBox()
		box6.Append(zero,false)


		mainBox:=ui.NewVerticalBox()
		mainBox.Append(box1,false)
		mainBox.Append(box2,false)
		mainBox.Append(box3,false)
		mainBox.Append(box4,false)
		mainBox.Append(box5,false)
		mainBox.Append(box6,false)

		window:=ui.NewWindow("Calculater",300,350,false)
		window.SetChild(mainBox)
		window.SetMargined(true)
		window.OnClosing(func(window *ui.Window) bool {
			ui.Quit()
			return true
		})

		seven.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"7")
		})
		eight.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"8")
		})
		nine.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"9")
		})
		four.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"4")
		})
		five.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"5")
		})
		six.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"6")
		})
		one.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"1")
		})
		two.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"2")
		})
		three.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"3")
		})
		zero.OnClicked(func(button *ui.Button) {

			var cur string
			cur=input1.Text()

			input1.SetText(cur+"0")
		})

		ts.OnClicked(func(button *ui.Button) {
			input1.SetText("")

			ttl=0
		})

		plus.OnClicked(func(button *ui.Button) {
			val=input1.Text()
			it,err:=strconv.Atoi(val)
			if (ttl==0) {
				ttl=it
				result=ttl
				input1.SetText("")
			}
			if(err==nil) {
				fmt.Println(it)
			}
			count=1
			input1.SetText("")
		})
		minus.OnClicked(func(button *ui.Button) {
			val=input1.Text()
			it,err:=strconv.Atoi(val)

			if (ttl==0) {
				ttl = it
				result = ttl
				input1.SetText("")
			}
			if(err==nil) {
				fmt.Println(it)
			}
			count=2
			input1.SetText("")
		})
		mul.OnClicked(func(button *ui.Button) {
			val=input1.Text()
			it,err:=strconv.Atoi(val)
			if (ttl==0) {
				ttl = it
				result=ttl
				input1.SetText("")
			}

			if(err==nil) {
				fmt.Println(it)
			}
			count=3
			input1.SetText("")

		})
		divide.OnClicked(func(button *ui.Button) {
			val=input1.Text()
			it,err:=strconv.ParseFloat(val,64)
			if (dv==0) {
				dv = it
				fresult=dv
				input1.SetText("")
			}
			//}else {
			//	it=dv/it
			//	at:=fmt.Sprintf("%f",it)
			//	//answer.SetText(at)
			//	fmt.Println(it)
			//	input1.SetText(at)
			//}

			if (err==nil) {
				fmt.Println("------")
			}
			count=4
			input1.SetText("")
		})
		
		equl.OnClicked(func(button *ui.Button) {
			switch count {
			case 1:
				var z=input1.Text()
				it,err:=strconv.Atoi(z)
				result=result+it
				var x string=strconv.Itoa(result)
				input1.SetText(x)
				if (err==nil) {
					fmt.Println("err",it)
				}
			case 2:
				var z=input1.Text()
				it,err:=strconv.Atoi(z)
				result=result-it
				var x string=strconv.Itoa(result)
				input1.SetText(x)
				if(err==nil) {
					fmt.Println("err",it)
				}
				fmt.Println(result)
			case 3:
				var x=input1.Text()
				it,err:=strconv.Atoi(x)

				result=result*it

				var z =strconv.Itoa(result)

				input1.SetText(z)

				if (err==nil) {
					fmt.Println("err",it)
				}
				fmt.Println(result)
			case 4:
				var a=input1.Text()

				it,err:=strconv.ParseFloat(a,64)
				fresult=fresult/it
				at:=fmt.Sprintf("%f",fresult)
				input1.SetText(at)

				if(err==nil){
					fmt.Println("err",it)
				}
				fmt.Println(fresult)

			}
		})
		window.Show()


	})
	if (myCal==nil) {
		fmt.Println("Null")
	}else {
		fmt.Println("Ok")
	}
}
