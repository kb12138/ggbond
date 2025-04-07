// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract MyTokenERC20 is IERC20 {
    // constructor() ERC20("MyToken", "MTK") ERC20Permit("MyToken") {

    // }

    string private name;
    string private symbol;
    uint256 private decimals;

    uint256 private _totalSupply;
    // 余额数据
    mapping(address => uint256) private balances;
    // 授权数据
    mapping(address account => mapping(address spender => uint256 value)) private _allowance;


    constructor(string memory name_, string memory symbol_) {
        name = name_;
        symbol = symbol_;
    }

    function getName() virtual external view  returns (string memory) {
        return name;
    }

    /**

    合约中的这五个方法都是读取的方法，
    name返回的是代币的名称，
    symbol返回的是代币的标志，
    decimals返回的是代币的精度，
    totalSupply返回的是代币的总数量，
    balanceOf返回到是某个账户的余额。
    */

    /**
     * @dev Returns the value of tokens in existence.
     */
    function totalSupply() virtual external view returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev Returns the value of tokens owned by `account`.
     * @dev 返回账户`account`所持有的代币数.
     */
    function balanceOf(address account)external view virtual returns (uint256)
    {
        return balances[account];
    }


    function transfer(address to, uint256 value) external returns (bool) {
        require(balances[msg.sender] >= value, "balance of sender is insufficient");

        transfer0(msg.sender, to, value);

        return true;
    }

    /**
    * 执行转账，并发送tranfer事件
    */
    function transfer0(address from, address to, uint256 value) internal {
        balances[from] -= value;
        balances[to] += value;

        emit Transfer(from, to, value);
    }


    function allowance(address owner, address spender)
    external
    view
    returns (uint256)
    {
        return _allowance[owner][spender];
    }


    /**
    * 授权，调用者为owner
    */
    function approve(address spender, uint256 value) external returns (bool) {
        address owner = msg.sender;
        require(balances[owner] >= value, "owner balance is insufficient");

        approve0(owner, spender, value);
        return true;
    }

    function approve0(address owner, address spender, uint256 value) internal  {
        _allowance[owner][spender] = value;
        emit Approval(owner, spender, value);
    }


    /**
    * 转账，调用者为spender， from -> to
    */
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) external returns (bool) {
        address spender = msg.sender;
        // 校验授权资金是否足够转账
        require(_allowance[from][spender] >= value, "allowance of spender exceeds allowed");

        // 扣除授权金额
        approve0(from, spender, _allowance[from][spender] - value);

        // 校验余额是否足够转账
        require(balances[from] >= value, "balance of owner exceeds allowed");
        transfer0(from, to, value);

        return true;
    }

    // 铸币， 新增资金， 总供给增加
    function mint(address account, uint256 value) external  returns (bool) {
        _totalSupply += value;
        balances[account] += value;

        emit Transfer(address(0), account, value);
        return true;
    }

    // 销毁 减少资金，总供给减少
    function burn(address account, uint256 value) external  returns (bool) {
        _totalSupply -= value;
        balances[account] -= value;

        emit Transfer(account, address(0), value);
        return true;
    }


}
