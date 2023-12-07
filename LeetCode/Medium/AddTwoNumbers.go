import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
 }


func print(a ...interface{}) {
    fmt.Println(a...)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var out *ListNode = new(ListNode)
    var miniSum int
    var c int = 0
    var f1, f2 bool
    var temp,temp1, temp2 *ListNode = out, l1, l2
    for !f1 || !f2{
        miniSum = 0
        if !f1 {
            miniSum += temp1.Val
            if temp1.Next != nil{
                temp1 = temp1.Next
            }else{
                f1 = true
            }
        }
        if !f2 {
            miniSum += temp2.Val
            if temp2.Next != nil{
                temp2 = temp2.Next
            }else{
                f2 = true
            }
        }
        miniSum += c
        c = 0
        temp.Val = miniSum % 10
        c = miniSum / 10
        if !f1 || !f2{
        temp.Next = new(ListNode)
        temp = temp.Next
        }
    }
    if c!= 0{
        temp.Next = new(ListNode)
        temp.Next.Val = c
    }
    return out
}
