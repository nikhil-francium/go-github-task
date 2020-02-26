
This project is about calculating the number of commits day wise in a week using week count, repo details and sort order.

***

# How to execute this project

***Input:***

  `weeks` : How many weeks into the past to consider when calculating the per-day average commits, default to 1 year (52 weeks)   `Eg: 52`
  
  `repo` : Input a repository name  `Eg: golang/tour`
  
  `sort` : Display order (ascending/descending) `Eg: asc/desc`
  
***  

***Command:*** `go run main.go github.go input.go -weeks={{weeks}} -repo={{repo}} -sort={{sort}}`

***Sample Command:*** `go run main.go github.go input.go -weeks=52 -repo=golang/tour -sort=desc`


This will display the commit count day wise in a week in the given sort order.

***

***Output:***

`Monday : 3
Tuesday : 3
Wednesday : 3
Thursday : 2
Friday : 1
Saturday : 1
Sunday : 0`
