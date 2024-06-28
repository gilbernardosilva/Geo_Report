export default function generatePassword(length = 12) {
  const lowercase = "abcdefghijklmnopqrstuvwxyz";
  const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  const numbers = "0123456789";
  const special = "@$!%*?&";
  const allChars = lowercase + uppercase + numbers + special;

  const randomValues = new Uint32Array(length); // Create array for random values
  window.crypto.getRandomValues(randomValues); // Fill with random values

  let password = "";

  password += lowercase.charAt(randomValues[0] % lowercase.length);
  password += uppercase.charAt(randomValues[1] % uppercase.length);
  password += numbers.charAt(randomValues[2] % numbers.length);
  password += special.charAt(randomValues[3] % special.length);

  for (let i = 4; i < length; i++) {
    password += allChars.charAt(randomValues[i] % allChars.length);
  }

  password = password
    .split("")
    .sort(() => {
      return 0.5 - crypto.getRandomValues(new Uint32Array(1))[0] / 4294967295;
    })
    .join("");

  return password;
}