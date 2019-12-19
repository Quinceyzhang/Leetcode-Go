package main

/*
	该例子中，主要采用蔡勒公式进行解答会比较简单
	公式的形式如下：
			//w=(d+2*m+3*(m+1)/5+y+y/4-y/100+y/400+1)%7
			w：周
			m：月
			y：年
			d:天
		公式要求month至少大于3，若为1或2，则将其作为上一年的13月和14月。
*/
func main() {
	//var year, month, day int

	println(dayOfTheWeek(29, 2, 2016))
}
func dayOfTheWeek(day int, month int, year int) string {
	var weeks = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	if month == 1 || month == 2 {
		year--
		month = month + 12
	}

	//var week = yearNum + (yearNum / 4) + (century / 4) - 2*century + (26 * (month + 1) / 10) + day - 1
	var week = day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1
	//w=(d+2*m+3*(m+1)/5+y+y/4-y/100+y/400+1)%7

	var index = week % 7

	var date = weeks[index]
	return date

}
