package basics

import "fmt"


func main() {
	var choice,num1,num2,result int
	fmt.Println("Enter num1")
	fmt.Scanf("%d",&num1)
	fmt.Println("Enter num2")
	fmt.Scanf("%d",&num2)
	for{
	fmt.Println("Enter which option do you want: \n 1.Add \n 2.Subtract \n 3.Multiplication \n 4.Remaninder \n 5.End")
	fmt.Scanf("%d",&choice)
	fmt.Println(choice)
	if choice ==1{
		result=num1+num2
		fmt.Println(result)
	}else if choice==2{
		result=num1-num2
		fmt.Println(result)
	}else if choice==3{
		result=num1*num2
		fmt.Println(result)
	}else if choice==4{
		result=num1/num2
		fmt.Println(result)
	}else if choice == 5{
		result=num1%num2
		fmt.Println(result)
	}else if choice==6{
		break
	}}

}
