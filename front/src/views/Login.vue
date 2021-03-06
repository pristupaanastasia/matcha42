<template>
  <div>
    <form >
      <p v-if="errors.length">
        <b>Пожалуйста исправьте указанные ошибки:</b>
      <ul>
        <li class="list-group-item" v-for="error in errors"  v-bind:key="error"></li>
      </ul>
      </p>
      <p>
        <label for="login">Login</label>
        <input
            id="login"
            v-model="login"
            type="text"
            name="login"
        >
      </p>

      <p>
        <label for="password">Password</label>
        <input
            id="password"
            v-model="password"
            type="text"
            name="password"
        >
      </p>


      <button @click='submit()'>Login</button>

    </form>
  </div>
</template>
<script>

const apiUrl = 'http://localhost:9000/api.user.login'


export default {
  el: '#app',
  data() {
    return {
      errors: [],
      login: null,
      password: null,
    }
  },

  methods: {
    checkForm: function () {
      this.errors = [];
      if (this.email && this.login && this.password && this.first_name && this.last_name) {

        return true;
      }
      if (!this.password) {
        this.errors.push('Требуется указать пароль.');
      }

      if (!this.login) {
        this.errors.push('Требуется указать логин.');
      }
    },

    async submit() {
      let s ={
        "login": this.login,

        "password":this.password,

      }
      await fetch(apiUrl,
          {
            method: "POST",
            credentials: "include",
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
            },
            cache: 'no-cache',
            redirect: 'follow',
            body: JSON.stringify(s)
          }).then(function(res){
        if (!res.ok) throw Error(`is not ok: ${res.status}`);
        //detect json
        if(res.headers.get("Content-Type").includes("json")){
          return res.json();
        }else{
          return res.text();
        }
      }).catch(function(res){
        //Errors
        console.log(res);
        //returned if errors
        return "Error!";
      });


    }
  },
};
</script>