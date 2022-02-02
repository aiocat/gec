#include <stack>
#include <iostream>
std::stack<int> stack;
int _rounded, _gec_one, _gec_two = 0;
class GecMath{
  public:
int factorial(int num){
int temp;
stack.push(num);
stack.push(num);
temp = stack.top();
stack.pop();
while(temp>1){
stack.push(temp);
stack.push(1);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
stack.push(_gec_two-_gec_one);
temp = stack.top();
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
stack.push(_gec_one*_gec_two);
};
return 0;
};
};
class GecString{
  public:
int print(){
while(stack.top()!=0){
_gec_one = stack.top();
stack.pop();
std::cout << (char)_gec_one;
};
stack.pop();
return 0;
};
int new_line(){
std::cout << (char)10;
std::cout << (char)13;
return 0;
};
int count(){
int temp;
stack.push(0);
temp = stack.top();
stack.pop();
while(stack.top()!=0){
stack.push(temp);
stack.push(1);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
stack.push(_gec_one+_gec_two);
temp = stack.top();
stack.pop();
stack.pop();
};
stack.pop();
stack.push(temp);
return 0;
};
};
int main(){
GecMath gmath;
GecString gstr;
stack.push(0);
stack.push(32);
stack.push(58);
stack.push(101);
stack.push(117);
stack.push(110);
stack.push(105);
stack.push(116);
stack.push(110);
stack.push(111);
stack.push(99);
stack.push(32);
stack.push(111);
stack.push(116);
stack.push(32);
stack.push(114);
stack.push(101);
stack.push(98);
stack.push(109);
stack.push(117);
stack.push(110);
stack.push(32);
stack.push(97);
stack.push(32);
stack.push(114);
stack.push(101);
stack.push(116);
stack.push(110);
stack.push(101);
stack.push(32);
stack.push(101);
stack.push(115);
stack.push(97);
stack.push(101);
stack.push(108);
stack.push(80);
_gec_one=gstr.print();
if(_gec_one!=0){return _gec_one;}
std::cin >> _gec_one;
stack.push(_gec_one);
_gec_one=gmath.factorial(stack.top());
if(_gec_one!=0){return _gec_one;}
std::cout << stack.top();
return 0;
};
