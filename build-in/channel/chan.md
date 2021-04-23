- buffered channel
- unbuffered channel
![img.png](../img/img.png)
  
```c++
if ((Q->rear + 1) % MAX_QSIZE == Q->front) //判满条件
```
  
![img.png](../img/img_1.png)
栈上的指针指向堆上的channel
