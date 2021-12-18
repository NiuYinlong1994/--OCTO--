//统计3个班的成绩情况，每个班有5名同学，求出各班的总成绩、平均分
//统计3个班级的总成绩、平均分（学生的成绩从键盘输入）
//统计3个班及格人数，每个班5个同学

package main

import (
	"fmt"
)

func main() {
	var classNum, studentNum int = 3, 5 //定义班级的数量和单个班级内学生的总人数，
	var sumAll float64
	var personCount int
	for class := 1; class <= classNum; class++ { //遍历所有班级
		var score, avg float64                               //定义变量学生分数和班级平均分 float64类型
		var sum float64 = 0.0                                //定义班级学生成绩总和初始化变量为0；float64类型
		for student := 1; student <= studentNum; student++ { //遍历一个班级的所有学生的分数
			fmt.Printf("请输入%d班第%d名同学的成绩: \n", class, student)
			fmt.Scanf("%v\n", &score)
			sum += score
			if score >= 60 { //判断班级内及格同学
				personCount++ //如果大于等于60，及格人数加1
			}
		}
		avg = sum / float64(studentNum)
		fmt.Printf("%v班的总成绩为:%v; 平均分为:%v \n", class, sum, avg)
		sumAll += sum
	}
	fmt.Printf("所有班级的总成绩为: %v; 平均分为: %v; 及格人数:%v\n", sumAll, sumAll/(float64(classNum)*float64(studentNum)), personCount)
}
