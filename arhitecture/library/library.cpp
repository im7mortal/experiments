//
// Created by petr on 16.02.17.
//
#include <iostream>

#include <mutex>
#include <condition_variable>
#include <chrono>
#include <thread>
#include <future>

class Calculator
{
    std::mutex _mutex;
    std::condition_variable cv;
    static Calculator *s_instance;

    bool finished;

    // singleton
    static Calculator *singleton;

    Calculator()
    {

    }
public:
    int get_value()
    {
        _mutex.lock();
        std::this_thread::sleep_for (std::chrono::seconds(1));
        _mutex.unlock();
        return 1;
    }
    void set_value(int v)
    {
        _mutex.lock();
        // logic
        std::cout << "thread" << std::endl;
        std::this_thread::sleep_for (std::chrono::seconds(1));
        _mutex.unlock();
    }
    void init(void)
    {
    }
   /* void inside(std::future<bool> fut)
    {
        std::cout << "inside" << std::endl;
    }*/
    static Calculator *instance()
    {
        if (!s_instance)
            s_instance = new Calculator();
        return s_instance;
    }
};

// Allocating and initializing GlobalClass's
// static data member.  The pointer is being
// allocated - not the object inself.
Calculator *Calculator::s_instance = 0;

bool is_prime (int x) {
  std::cout << "Calculating. Please, wait...\n";
        std::this_thread::sleep_for (std::chrono::seconds(5));
  for (int i=2; i<x; ++i) if (x%i==0) return false;
  return true;
}

extern "C" void set(int i);
void set(int i)
{
    std::future<bool> fut = std::async (is_prime,313222313);
        std::cout << i << std::endl;
    Calculator::instance()->set_value(0);
//    Calculator::instance()->inside(fut);
        std::cout << "Wait async" << std::endl;
    fut.get();
    return ;
}

extern "C" int get(void);
int get(void)
{
    Calculator::instance()->set_value(0);
}

extern "C" int init(void);
int init(void) {
    return 0;
}
