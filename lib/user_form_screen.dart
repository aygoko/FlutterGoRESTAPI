import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:flutter_keyboard_visibility/flutter_keyboard_visibility.dart'; 
import 'user_bloc.dart';
import 'user_model.dart';

class UserFormScreen extends StatefulWidget {
  @override
  _UserFormScreenState createState() => _UserFormScreenState();
}

class _UserFormScreenState extends State<UserFormScreen> {
  final _formKey = GlobalKey<FormState>();
  final _loginController = TextEditingController();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();
  bool _isLoading = false;
  bool _hasError = false;
  String _errorMessage = '';
  final _keyboardVisibilityController = KeyboardVisibilityController(); 

  Future<void> _submitForm() async {
    if (!_formKey.currentState!.validate()) return;

    setState(() {
      _isLoading = true;
      _hasError = false;
    });

    try {
      final response = await http.post(
        Uri.parse('http://10.0.2.2:8080/api/users'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({
          'login': _loginController.text.trim(),
          'email': _emailController.text.trim(),
          'password': _passwordController.text.trim(),
        }),
      );

      if (response.statusCode == 201) {
        Navigator.pop(context, {'refresh': true});
      } else {
        final errorData = json.decode(response.body);
        setState(() {
          _hasError = true;
          _errorMessage = errorData['error'] ?? 'Server returned ${response.statusCode}';
        });
      }
    } catch (e) {
      setState(() {
        _hasError = true;
        _errorMessage = 'Connection error: ${e.toString()}';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Create User'),
      ),
      body: _isLoading 
          ? Center(child: CircularProgressIndicator())
          : Padding(
              padding: EdgeInsets.all(16),
              child: Form(
                key: _formKey,
                child: Column(
                  children: [
                    TextFormField(
                      controller: _loginController,
                      decoration: InputDecoration(
                        labelText: 'Login',
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Login is required';
                        }
                        return null;
                      },
                    ),
                    SizedBox(height: 16),
                    TextFormField(
                      controller: _emailController,
                      keyboardType: TextInputType.emailAddress,
                      textInputAction: TextInputAction.next,
                      decoration: InputDecoration(
                        labelText: 'Email',
                        suffixIcon: IconButton(
                          icon: Icon(Icons.clear),
                          onPressed: () => _emailController.clear(),
                        ),
                        border: OutlineInputBorder(),
                      ),
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Email is required';
                        }
                        if (!value.contains('@')) {
                          return 'Invalid email format';
                        }
                        return null;
                      },
                    ),
                    SizedBox(height: 16),
                    TextFormField(
                      controller: _passwordController,
                      decoration: InputDecoration(
                        labelText: 'Password',
                        border: OutlineInputBorder(),
                      ),
                      obscureText: true,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Password is required';
                        }
                        if (value.length < 6) {
                          return 'Minimum 6 characters';
                        }
                        return null;
                      },
                    ),
                    SizedBox(height: 24),
                    ElevatedButton(
                      onPressed: _isLoading ? null : _submitForm,
                      child: _isLoading 
                        ? CircularProgressIndicator(color: Colors.white,)
                        : Text('Create User'),
                      style: ElevatedButton.styleFrom(
                        minimumSize: Size(double.infinity, 48),
                        backgroundColor: Colors.blue,
                      ),
                    ),
                    if (_hasError)
                      Padding(
                        padding: EdgeInsets.only(top: 16),
                        child: Text(
                          _errorMessage,
                          style: TextStyle(color: Colors.red, fontSize: 14),
                        ),
                      ),
                  ],
                ),
              ),
            ),
    );
  }

  @override
  void dispose() {
    _loginController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }
}