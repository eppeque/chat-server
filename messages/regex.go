package messages

const RxRegister = `^REGISTER ([a-zA-Z0-9_]{5,}) \$2b\$12\$([a-zA-Z0-9.\/]{22})([a-zA-Z0-9.\/]{31})$`
const RxLogin = `^LOGIN ([a-zA-Z0-9_]{5,})$`
const RxConfirm = `^CONFIRM ([a-z0-9]{64})$`
