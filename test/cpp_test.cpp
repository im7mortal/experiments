#include <cglog.h>
#include<cstring>
#include <string>
using namespace std;

extern "C" void TestCGLOG();
void TestCGLOG() {
    string str = "Log from C++";
    char *cstr = new char[str.length() + 1];
    strcpy(cstr, str.c_str());
    Info(cstr);
    delete [] cstr;
}
