#include <thread>
#include <iostream>
#include <chrono>

void t1_func()
{
    std::chrono::duration<int, std::milli> dur(2000);
    std::this_thread::sleep_for(dur);
    std::cout << "t1 says hello" << std::endl;
}

void t2_func()
{
    std::cout << "t2 says hello" << std::endl;
}

int main()
{
    std::thread t1(t1_func);
    std::thread t2(t2_func);

    t1.join();

    t2.join();
}
