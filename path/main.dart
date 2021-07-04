import 'dart:io';

void main(List<String> args) {
  print("hello WOrld");
  File('./main.dart').readAsString().then((value) => print(value));
}
