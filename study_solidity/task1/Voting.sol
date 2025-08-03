// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    mapping(string => uint) internal userName2VotingMap;
    string[] internal userNameKeys;

    function vote(string memory userName) public {
        uint voting = userName2VotingMap[userName];
        if (voting == 0) {
            userNameKeys.push(userName);
        }
        voting++;
        userName2VotingMap[userName] = voting;
    }

    function getVotes(string memory userName) public returns (uint voting) {
        voting = userName2VotingMap[userName];
    }

    function resetVotes() public {
        for (uint idx = 0; idx < userNameKeys.length; idx++) {
            userName2VotingMap[userNameKeys[idx]] = 0;
        }
    }
}
