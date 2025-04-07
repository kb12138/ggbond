// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract MyTokenERC721 is IERC721 {
    // Token name
    string private _name;

    // Token symbol 铸币代号
    string private _symbol;

    mapping(uint256 tokenId  => address) private _owners;

    mapping(address owner => uint256) private _balances;

    // 单个token授权
    mapping(uint256 tokenId => address) private _tokenApprovals;
    // 批量授权
    mapping(address owner => mapping(address => bool)) private _operatorApprovals;

    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
    }

    function balanceOf(address owner) external view override returns (uint256) {
        require(owner != address(0), "owner is zero");
        return  _balances[owner];
    }

    function ownerOf(uint256 tokenId) external view override returns (address) {
        return _owners[tokenId];
    }

    function supportsInterface(bytes4 interfaceId) external pure   returns (bool) {
        require((interfaceId == type(IERC721).interfaceId
        // || super.supportsInterface(interfaceId)
        ), "interfaceID not supported");
        return true;
    }

    // 铸币
    function _mint(address to, uint256 tokenId) internal {
        require(to != address(0), "to is zero;");
        // 更新余额
        _balances[to] = _balances[to] +1;
        // 更新owner
        _owners[tokenId] = to;
        // 更新 approve
        _tokenApprovals[tokenId] = to;
        // 发送消息
        emit Transfer(address(0), to, tokenId);
        emit Approval(address(0), to, tokenId);
    }

    // 销毁
    function _burn(uint256 tokenId) internal {
        address owner = _owners[tokenId];

        // 更新余额
        _balances[owner] = _balances[owner] - 1;
        // 更新owner
        _owners[tokenId] = address(0);
        // 更新 approve
        _tokenApprovals[tokenId] = address(0);
        // 发送消息
        emit Transfer(owner,address(0), tokenId);
        emit Approval(owner, address(0), tokenId);
    }


    function safeTransferFrom(address from, address to, uint256 tokenId, bytes calldata data) public virtual   {
        safeTransferFrom0(from, to, tokenId, data);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public  {
        safeTransferFrom0(from, to, tokenId, "");
    }

    function safeTransferFrom0(address from, address to, uint256 tokenId, bytes memory data) internal    {
        transferFrom(from, to, tokenId);
        ERC721Utils.checkOnERC721Received(msg.sender, from, to, tokenId, data);
    }


    // called by from
    function transferFrom(address from, address to, uint256 tokenId) public  {
        require(to != address(0), "to is zero");
        require(_owners[tokenId] == from,"from is not owners");
        //  require (_balances[to] + _balances[_owners[tokenId]] >= _balances[to], "balance overflow");
        // 更新余额
        _balances[from] = _balances[from] - 1;
        _balances[to] = _balances[to] + 1;
        // 更新owner
        _owners[tokenId] = to;
        // 更新 approve
        _tokenApprovals[tokenId] = to;
        // 发送消息
        emit Transfer(from, to, tokenId);
        emit Approval(from, to, tokenId);

    }

    function approve(address to, uint256 tokenId) external override {
        // check owner
        address owner = msg.sender;
        require(_owners[tokenId] == owner,"owner is not owners");
        _tokenApprovals[tokenId] = to;
        emit Approval(owner,to,tokenId);
    }


    //  called by owner
    function setApprovalForAll(address operator, bool approved) external{
        require(operator != address(0),"operator is zero");

        address owner = msg.sender;
        _operatorApprovals[owner][operator] = approved;
        emit ApprovalForAll(owner, operator, approved);
    }


    function getApproved(uint256 tokenId) virtual override external view returns (address operator){
        return _tokenApprovals[tokenId];
    }


    function isApprovedForAll(address owner, address operator) external view returns (bool){
        return _operatorApprovals[owner][operator];
    }

}
