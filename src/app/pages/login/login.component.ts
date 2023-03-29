import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from 'app/auth.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
    form: FormGroup;
    errorMessage: string = '';

    constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
        this.form = this.fb.group({
            // sign up info
            email: ['', Validators.required],
            signupUsername: ['', Validators.required],
            signupPassword: ['', Validators.required],
            // sign in info
            signinUsername: ['', Validators.required],
            signinPassword: ['', Validators.required],
        });
    }

    ngOnInit() { }

    login() {
        const val = this.form.value;

        if (val.signinUsername && val.signinPassword) {
            this.authService.login(val.signinUsername, val.signinPassword, val.form);
            if (this.router.url != 'http://localhost:4200/') {
                this.errorMessage = 'Invalid credentials. Please try again.';
            }
        }
    }

    signup() {
        const vals = this.form.value;

        if (vals.email && vals.signupUsername && vals.signupPassword) {
            this.authService.signup(vals.email, vals.signupUsername, vals.signupPassword);
        }
    }
}