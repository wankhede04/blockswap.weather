// SPDX-License-Identifier: MIT
pragma solidity =0.8.18;

import { IRegistration } from "./interfaces/IRegistration.sol";

/**
 * @title Registration Contract
 * @notice This contract is used to perform registration, and resign.
 * @dev Registered participants can call the server maximum once every 12 seconds to report the weather.
 */
contract Registration is IRegistration {
    // Stores participants status
    mapping(address => LifecycleStatus) public participants;

    /// @inheritdoc IRegistration
    function register() external virtual {
        require(participants[msg.sender] == LifecycleStatus.Unregistered, "Already registered or resigned");

        participants[msg.sender] = LifecycleStatus.Registered;
        emit ParticipantRegistered(msg.sender);
    }

    /// @inheritdoc IRegistration
    function resign() external virtual {
        require(participants[msg.sender] == LifecycleStatus.Registered, "Not registered");

        participants[msg.sender] = LifecycleStatus.Resigned;
        emit ParticipantResigned(msg.sender);
    }
}
