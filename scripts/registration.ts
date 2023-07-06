import { ethers, run } from "hardhat";

const registration = async () => {
  const Registration = await ethers.getContractFactory("Registration");

  console.log("Deploying ...");
  const registration = await Registration.deploy();
  await registration.deployed();
  console.log("Deployed: ", registration.address);

  try {
    await run("verify:verify", {
      address: registration.address,
    });
  } catch (error) {
    console.log("ERROR - verify - perp view!");
  }
};

registration()
  .then(() => process.exit(0))
  .catch(error => {
    console.error("Error: ", error);
    process.exit(1);
  });
