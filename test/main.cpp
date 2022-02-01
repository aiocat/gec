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
int main(){
GecMath gmath;
_gec_one=gmath.factorial(10);
if(_gec_one!=0){return _gec_one;}
_gec_one = stack.top();
stack.pop();
std::cout << _gec_one;
};
