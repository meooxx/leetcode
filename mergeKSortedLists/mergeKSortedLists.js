


class ListNode {
  constructor(val = null, next = null) {
    this.val = val 
    this.next = next
  }
}




const mergeKLists = (lists) => {
  var flag = true
  var result = null
  var r = new ListNode()
  const len = lists.length
  for(;flag;) {
    var count = 0  
    var temp = -1
    var minList

    if(len === 0) return null
    
    lists.forEach((list, index)=> {
      if(list != undefined &&
          (temp == -1 || list.val < lists[temp].val) 
        ){
          temp = index
          minList =  lists[temp]
        } 
      if( index === len -1 ) {
        if(temp != -1) {
          r.val = minList.val
          lists[temp] = minList.next
          if(lists[temp] == undefined  && index !== temp ) {
            count ++
          }
        }
      }
      if(lists[index] == undefined) {
        count ++
      }
      if(count === len) {
        flag = false
      }
    })

    if(result == null && r.val != undefined) {
      result = r
    }
    if(flag) {
      r.next = new ListNode()
      r = r.next
    }

  }
  return result
}


const l1 = new ListNode(1)
const l2 = new ListNode(2)
const l3 = new ListNode(3)

l1.next = l2
l2.next = l3

const n1 = new ListNode(1)
const n2 = new ListNode(2)
const n3 = new ListNode(3)

n1.next = n2
n2.next = n3

const readline = require('readline')

const line = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  prompt: "input add  start:  "
})

line.prompt()

line.on("line", (input) => {
  console.log("start")
  var sortedList= mergeKLists([l1, n1, l1])
  while(sortedList) {
    console.log(sortedList)
    sortedList = sortedList.next
  }
})
.on("close", () => {
  console.log("end")
  process.exit(0)
})



var sortedList = mergeKLists([l1, n1, l1])
/* while(sortedList) {
  console.log(sortedList)
  sortedList = sortedList.next
} */