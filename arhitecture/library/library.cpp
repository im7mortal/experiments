//
// Created by petr on 16.02.17.
//
#include <iostream>

#include <mutex>
#include <condition_variable>
#include <chrono>
#include <thread>
#include <future>
#include <functional>
#include <map>

struct response {
  int length;
  int data[3];
};

class Calculator
{
    std::mutex _mutex;
    std::condition_variable cv;
    std::map<int,response> global_queue;

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
   void inside(std::future<response>& fut)
    {
        response res = fut.get();
        std::cout << res.data[0] << std::endl;
    }
   void inst(response res)
    {
        std::cout << "Response" << std::endl;
        std::cout << res.data[0] << std::endl;
        std::cout << "Response" << std::endl;
        global_queue[res.length] = res;
    }
    response inst1(int query)
    {
        bool found = true;
        while (found){
             if ( global_queue.find(query) == global_queue.end() ) {
              // not found
               std::this_thread::sleep_for (std::chrono::seconds(5));
            } else {
             return global_queue[query];

            }
        }

    }

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
response is_prime2 (int x) {
    return Calculator::instance()->inst1(x);
}

extern "C" void set(int i);
void set(int i)
{
    std::future<response> fut = std::async (is_prime2,i);
        std::cout << i << std::endl;
    Calculator::instance()->set_value(0);
    Calculator::instance()->inside(std::ref(fut));

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

extern "C" void sent(int i[], int j, int k);
void sent(int i[], int j, int k) {
        std::cout << "SENT!" << std::endl;
        int lol[3];
        response res;
        res.length = k;

        res.data[0] =  i[0];
        res.data[1] =  i[1];
        res.data[2] =  i[2];

        Calculator::instance()->inst(res);
    return ;
}
