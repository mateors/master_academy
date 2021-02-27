package main

import "fmt"

type Email interface{

 Write(string, string, string)
 Send() string
 Read()

}

type Test struct{
 To string
 From string
 Subject string
 Message string
}

func (t Test) Write(to string, from string, subject string){
 fmt.Println(to, from, subject)
}

func (t Test) Send() string{
 //fmt.Println(t.To, "email sent")
 return t.To
}

func (t Test) Read(){
 fmt.Println(t.From, "received from")
}


type Doctor struct{
  Name string
  Education string
  Age int
  Salary float32
}

//method
func (d Doctor)Speak(){
 fmt.Println(d.Name,"can speak")
}

//method
func (d Doctor)getName() string{
 return d.Name
}

//method
func (d Doctor)getSalaryInfo() float32{
 return d.Salary
}

func main(){

 //Literals
 //var d =Doctor{ "Tareq", "MBBS", 30, 50000.00, }
 //var d Doctor{ Name:"Tareq", 
	//Education:"MBBS", Age:30, Salary:50000.00, }

 //var d=Doctor{  Education:"MBBS"
  //	 Age:30, Salary:50000.00, Name:"Tareq", }

 /*
  var d Doctor
  d.Name="Tareq"
  d.Education="BDS"
  d.Age=30
  d.Salary=50000.50
  
  fmt.Println(d.getName())
  fmt.Println(d.getSalaryInfo())

*/

 //var e Email
 //fmt.Println(e)

 var tst Test
 tst.To="billahmdmostain@gmail.com"
tst.From="admin@mateors.com"
tst.Subject="Test email"
tst.Message="Hello this is a test email"

tst.Write(tst.To, tst.From, tst.Subject)
 
 
}