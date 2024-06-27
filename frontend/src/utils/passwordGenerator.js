export default function generatePassword(length = 12) {
    const lowercase = "abcdefghijklmnopqrstuvwxyz";
    const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    const numbers = "0123456789";
    const special = "@$!%*?&";
    const allChars = lowercase + uppercase + numbers + special;

    let password = "";

    password += lowercase.charAt(Math.floor(Math.random() * lowercase.length));
    password += uppercase.charAt(Math.floor(Math.random() * uppercase.length));
    password += numbers.charAt(Math.floor(Math.random() * numbers.length));
    password += special.charAt(Math.floor(Math.random() * special.length));

    for (let i = 4; i < length; i++) {
      password += allChars.charAt(Math.floor(Math.random() * allChars.length));
    }

    password = password.split('').sort(function(){return 0.5-Math.random()}).join('');

    return password;
  }