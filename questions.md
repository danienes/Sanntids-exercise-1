Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Your answer here*
Parallellism is several things happening at once, while concurrency is more like several things is using one compiler for example at the same time, however they are not executed at the same time. 

What is the difference between a *race condition* and a *data race*? 
> *Your answer here* 
A race condition is a logical error where the result depends on ordering in time. A data race is a type of race condition which is about unsecured access to shared memory which can lead to errors. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *Your answer here* 
It decides what thread should run next. It does it by calling a yield/reschedule function. 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Your answer here*
It solves collisions, for example errors which can occur when several things try to change a variable for instance at the same time. 

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*
They are mechanisms to run several tasks at the same time. They are simpler than normal threads, and have a lower cost. 

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*
Both, it is harder because it can be pretty advanced and you have to be careful, however it opens a lot of doors and improves the performance.


What do you think is best - *shared variables* or *message passing*?
> *Your answer here*
Message passing because it reduces the risk of race conditions and is less complex than shared variables.

