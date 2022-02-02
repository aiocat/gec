#include <stack>
#include <iostream>
#include <vector>
std::stack<int> stack_0;
int _rounded, _gec_one, _gec_two = 0;
class GecMath{
  public:
int factorial(int num){
int temp;
stack_0.push(num);
stack_0.push(num);
temp = stack_0.top();
stack_0.pop();
while(temp>1){
stack_0.push(temp);
stack_0.push(1);
_gec_one = stack_0.top();
stack_0.pop();
_gec_two = stack_0.top();
stack_0.pop();
stack_0.push(_gec_two-_gec_one);
temp = stack_0.top();
_gec_one = stack_0.top();
stack_0.pop();
_gec_two = stack_0.top();
stack_0.pop();
stack_0.push(_gec_one*_gec_two);
};
return 0;
};
};
class GecString{
  public:
int print(){
while(stack_0.top() != 0){
_gec_one = stack_0.top();
stack_0.pop();
std::cout << (char)_gec_one;
};
stack_0.pop();
return 0;
};
int new_line(){
std::cout << (char)10;
std::cout << (char)13;
return 0;
};
int count(){
int temp;
stack_0.push(0);
temp = stack_0.top();
stack_0.pop();
while(stack_0.top() != 0){
stack_0.push(temp);
stack_0.push(1);
_gec_one = stack_0.top();
stack_0.pop();
_gec_two = stack_0.top();
stack_0.pop();
stack_0.push(_gec_one+_gec_two);
temp = stack_0.top();
stack_0.pop();
stack_0.pop();
};
stack_0.pop();
stack_0.push(temp);
return 0;
};
};
int main(){
GecMath gmath;
GecString gstr;
stack_0.push(0);
stack_0.push(99);
stack_0.push(98);
stack_0.push(97);
_gec_one=gstr.print();
if(_gec_one!=0){return _gec_one;}
_gec_one=gstr.new_line();
if(_gec_one!=0){return _gec_one;}
_gec_one=gmath.factorial(5);
if(_gec_one!=0){return _gec_one;}
_gec_one = stack_0.top();
stack_0.pop();
std::cout << _gec_one;
return 0;
};
