# Find Symmetric Poems

![](https://github.com/LingDong-/magic-square-poems/raw/main/illustration.svg)

I got inspirations from https://github.com/LingDong-/magic-square-poems .

"Observe the above 5x5 matrix of Chinese characters: when read horizontally, it consists of five lines from five different poems, and so does it when read vertically, of the exact same lines. "

I prefer symmetric matrix but not magic matrix, because I think it's a symmetric matrix.
Observe that when a matrix is symmetric, as in these cases, the matrix is equal to its transpose.

To solve this problem, I implemented a quicker parallel alogorithm. It only takes 36 seconds to find the symmetric poems.

*Run*
```
go run main.go
```

*Output*
```
❯ go run main.go               
风月清江夜
月出夜山深
清夜方归来
江山归谢客
夜深来客稀

心如七十人
如何十年间
七十未成事
十年成底事
人间事事慵

[[风月清江夜 月出夜山深 清夜方归来 江山归谢客 夜深来客稀] [心如七十人 如何十年间 七十未成事 十年成底事 人间事事慵]]

~/go/src/FindSymmetricPoems 36s
```
