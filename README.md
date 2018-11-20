# GoPlayGround

This is an example project I used to learn some Go. It implements a main function using two implementations of the same interface. Essentially, a bridge pattern and a factory pattern.
The first implementation runs directly on the main thread. The second implementation uses a REST service. I tried to implement helper functions to generalize the implementation and reduce boilerplate code. Had to use reflection on the way. Oh well.
