package main

import "fmt"

func update(a *int, name *string) {

	*a = *a + 5               // defrencing pointer address
	*name = *name + " Billah" // defrencing pointer address
	fmt.Println(a, name)
}

func main() {

	// var p *int //0xc000006028=0xc0000100a0
	// var x int  //0xc0000100a0=100 ভ্যারিয়েবল ডিক্লেয়ার করা
	// //d := &doctor{"Sanzida", "MBBS"}

	// fmt.Println(&p, &x)
	// fmt.Println(p, x) //zero values pointer=nil, integar= 0

	// //ভ্যারিয়েবল এর মধ্যে মান বসানো
	// x = 10
	// p = &x
	// *p = 100 //dereferencing p ভ্যারিয়েবল এর মধ্যে যে এড্রেসভ্যালু আছে তার ভ্যারিয়েবল এর মান

	// var p2 **int
	// p2 = &p
	// //*p2=p
	// //**p2==*p2->p->x=100

	// fmt.Println(x, p, *p, p2, *p2, **p2)

	x := 5
	name := "mostain"
	update(&x, &name)
	fmt.Println(&x, &name, x, name)

}

/*
data store in a memory
আইডেন্টিফায়ার= ভ্যারিয়েবল এর নাম
মেমরি এড্রেস = RAM এর এড্রেস হেক্সাডেসিমাল ভ্যালু (0xc000006028)
ভ্যালু

রেগুলার ভ্যারিয়েবল এর মধ্যে ভ্যালু হিসেবে dataType(int,string,float,bool,struct,map,array) এর ডাটা থাকে
পয়েন্টার ভ্যারিয়েবল এর মধ্যে ভ্যালু হিসেবে হেক্সাডেসিমেল টাইপ (0xc000006028) এর ডাটা থাকে

পয়েন্টার::
অন্য কোনো ভ্যারিয়েবল এর মেমরি এড্রেস স্টোর করতে ব্যবহার করা হয়
অর্থাৎ পয়েন্টার টাইপ এর ভ্যারিয়েবল এর মধ্যে মেমরি এড্রেস ভ্যালু হিসেবে থাকে
ডিফল্ট বা জিরো ভ্যালু হচ্ছে nil
nil ভ্যালু থাকা অবস্থায় পয়েন্টার ভ্যারিয়েবল এর মধ্যে কোনো ভ্যালু এসাইন করতে পারবেন না
(** এক ভারিয়াব্লেকে অন্য কোনো ভ্যারিয়েবল এর সাথে কানেক্ট করার জন্য পয়েন্টার  ব্যবহার করা হয়)

*/
