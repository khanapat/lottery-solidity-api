solc --abi --bin lottery.sol -o build

abigen --bin=./build/Lottery.bin --abi=./build/Lottery.abi --pkg=feature --out=./feature/lottery.go

abigen --bin=./build/ERC20.bin --abi=./build/ERC20.abi --pkg=coin --out=./coin/erc20.go