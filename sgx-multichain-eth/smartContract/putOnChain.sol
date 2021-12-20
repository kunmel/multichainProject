pragma solidity ^0.8.7;



contract PutOnChain {
	mapping (address => mapping(uint => Project)) examer;
	mapping (uint => address) projectIdInStore;
	
	
	//项目信息结构
	struct Project {
		uint id;
		string price;
		string label;
	}

	//添加项目信息上链
	function addProject(uint _id, string memory _price, string memory _label) public {
		Project memory project = Project(_id, _price, _label);
		examer[msg.sender][_id] = project;
		projectIdInStore[_id] = msg.sender;
	}

	//查询展示链上项目信息
	function getProject(uint _id) view public returns (uint, string memory, string memory) {
		Project memory project = examer[projectIdInStore[_id]][_id];
		return (project.id, project.price, project.label);
	}
}

