const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");


module.exports = buildModule("TokenSampleProject", (m) => {

//   const lock = m.contract("Lock", [unlockTime], {
//     value: lockedAmount,
//   });

  return  m.contract("Token");
});
