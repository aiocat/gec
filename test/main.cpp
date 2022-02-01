#include <stack>
#include <iostream>
std::stack<int> stack;
int _rounded, _gec_one, _gec_two = 0;
int test(int a){
stack.push(a);
stack.push(a);
stack.push(1);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
stack.push(_gec_one+_gec_two);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
stack.push(_gec_one*_gec_two);
return 0;
}
int main(){
_gec_one=test(2);
if(_gec_one!=0){return _gec_one;}
int thing = stack.top();
stack.push(10);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();if(_gec_one>_gec_two){
std::cout << thing;
}
stack.push(104);
stack.push(101);
stack.push(108);
stack.push(108);
stack.push(111);
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
return 0;
}
