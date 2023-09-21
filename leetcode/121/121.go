package main

func maxProfit(prices []int) int {
  mxProfit := 0
  minSoFar := 999999
  for _, price := range prices {
    if price < minSoFar {
      minSoFar = price
    }

    profit := price - minSoFar
    if profit > mxProfit {
      mxProfit = profit
    }
  }

  return mxProfit
}
