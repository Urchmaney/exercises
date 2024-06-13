const User = function (name, gender, age) {
  this.name = name;
  this.gender = gender
  this.age = age;
}

const users = [
  new User('Albert Fiati-Kumasenu', 'Male', 21),
  new User('Eileen Ackabah', 'Female', 27),
  new User('Papa Yaw Owusu-Ankoma', 'Male', 45),
  new User('Naydia Frempong', 'Female', 52),
  new User('Berth Aba T.', 'Female', 12),
  new User('Lawrence Poku', 'Male', 71),
  new User('Joel Funu', 'Male', 32),
  new User('Selma Adu Twumwaa', 'Female', 10),
]

const displayUserDetails = function () {
  let i = 0;

  for (i; i < users.length; i++) {
    console.log('>' + (i + 1) + '< ' + users[i].name + ' is a ' + users[i].age + ' old ' + users[i].gender);
  }

}

displayUserDetails();
