#include <stack>
#include <iostream>
std::stack<int> stack;
int _rounded, _gec_one, _gec_two = 0;
int test(int a,int b){
stack.push(a);
stack.push(b);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
if(_gec_two%_gec_one==0){
stack.push(_gec_two/_gec_one);
_rounded = 0;}else{
stack.push((int)(_gec_two/_gec_one));
_rounded = 1;}
return 0;
}
int main(){
stack.push(66);
int qweasd = stack.top();
stack.pop();
_gec_one=test(qweasd,5);
if(_gec_one!=0){return _gec_one;}
_gec_one = stack.top();
stack.pop();
std::cout << _gec_one;
std::cout << (char)10;
stack.push(_rounded);
_gec_one = stack.top();
stack.pop();
std::cout << _gec_one;
return 0;
}
