//
// Created by petr on 16.02.17.
//
#include <iostream>

#include <mutex>
#include <condition_variable>
#include <chrono>
#include <thread>

class Calculator
{
    std::mutex *_mutex;
    std::condition_variable cv;

    bool finished;

    // singleton
    static Calculator *singleton;

    // background manager
    void manager() {
        std::cout << "Manager thread launched\n";
        // lock guard guarantee that lock will be released after function will be finished
        std::lock_guard<std::mutex> lck (_mutex);
        // loop
        while (!finished) {
            // check global state
            std::this_thread::sleep_for (std::chrono::seconds(1));
            cv.wait(lck)
        };
    }
    // constructor
    Calculator()
    {
        // set vars const
        // initialization
        std::this_thread::sleep_for (std::chrono::seconds(n));
        std::cout << "Inited\n";

        // start thread
        std::thread manager_thread (manager);
        // place thread in background as demon
        manager_thread.detach()
    }
public:


    int get_value()
    {
        _mutex.lock();
        return m_value;
    }
    void set_value(int v)
    {
        _mutex.lock();
        // logic
        std::this_thread::sleep_for (std::chrono::seconds(n));
        m_value = v;
        _mutex.unlock();
    }
    void init(void)
    {
    }
    static Calculator *instance()
    {
        if (!singleton)
            singleton = new Calculator();
        return singleton;
    }
};

// Allocating and initializing Calculator's
// static data member.  The pointer is being
// allocated - not the object inself.
Calculator *Calculator::s_instance = 0;

extern "C" int get(void);
int get(void)
{
    return Calculator::instance()->get_value();
}

extern "C" void set(int i);
void set(int i)
{
    Calculator::instance()->set_value(i);
}

int init(void) {
    // for now simple initialization
    Calculator::instance()->init();
    return 0;
}
