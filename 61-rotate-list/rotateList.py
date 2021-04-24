class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def rotateRight(self, head, k: int):
    if(not head):
        return
    if(k == 0):
        return head
    startNode = head
    index = 0

    # circular link list and cal length
    while head.next != None:
        head = head.next
        index += 1
    head.next = startNode
    head = startNode
    length = index+1
    start = length - (k % length)

    # find start
    count = 0
    while count != start:
        count += 1
        head = head.next

    startNode = head

    while index > 0:
        head = head.next
        index -= 1
    head.next = None

    return startNode
