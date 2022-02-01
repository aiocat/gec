#include <stack>
#include <iostream>
std::stack<int> stack;
int _rounded, _gec_one, _gec_two = 0;
int main(){
stack.push(5);
stack.push(61);
_gec_one = stack.top();
stack.pop();
_gec_two = stack.top();
stack.pop();
if(_gec_one%_gec_two==0){
stack.push(_gec_one/_gec_two);
_rounded = 0;}else{
stack.push((int)(_gec_one/_gec_two));
_rounded = 1;}
stack.push(_rounded);
_gec_one = stack.top();
stack.pop();
if(_gec_one==0){
_gec_one = stack.top();
stack.pop();
std::cout << _gec_one;
}else{std::cout << (char)65;
}
}
