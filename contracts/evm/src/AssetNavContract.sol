/**
 *Submitted for verification at Etherscan.io on 2025-10-29
*/

// SPDX-License-Identifier: MIT
// File: keystone/IERC165.sol


pragma solidity ^0.8.0;

interface IERC165 {
	function supportsInterface(bytes4 interfaceId) external view returns (bool);
}



// File: keystone/IReceiver.sol


pragma solidity ^0.8.0;


interface IReceiver is IERC165 {
	function onReport(bytes calldata metadata, bytes calldata report) external;
}



// File: keystone/IReceiverTemplate.sol


pragma solidity ^0.8.0;



/// @title IReceiverTemplate - Abstract receiver with workflow validation and metadata decoding
abstract contract IReceiverTemplate is IReceiver {
	// Immutable expected values
	address public EXPECTED_AUTHOR;
	bytes10 public EXPECTED_WORKFLOW_NAME;

	// Custom errors
	error InvalidAuthor(address received, address expected);
	error InvalidWorkflowName(bytes10 received, bytes10 expected);

	constructor(address expectedAuthor, bytes10 expectedWorkflowName) {
		EXPECTED_AUTHOR = expectedAuthor;
		EXPECTED_WORKFLOW_NAME = expectedWorkflowName;
	}

	/// @inheritdoc IReceiver
	function onReport(bytes calldata metadata, bytes calldata report) external virtual override {
		(address workflowOwner, bytes10 workflowName) = _decodeMetadata(metadata);

		if (workflowOwner != EXPECTED_AUTHOR) {
			revert InvalidAuthor(workflowOwner, EXPECTED_AUTHOR);
		}
		if (workflowName != EXPECTED_WORKFLOW_NAME) {
			revert InvalidWorkflowName(workflowName, EXPECTED_WORKFLOW_NAME);
		}

		_processReport(report);
	}

	/// @notice Extracts the workflow name and the workflow owner from the metadata parameter of onReport
	/// @param metadata The metadata in bytes format
	/// @return workflowOwner The owner of the workflow
	/// @return workflowName  The name of the workflow
	function _decodeMetadata(bytes memory metadata) internal pure returns (address, bytes10) {
		address workflowOwner;
		bytes10 workflowName;
		// (first 32 bytes contain length of the byte array)
		// workflow_id             // offset 32, size 32
		// workflow_name            // offset 64, size 10
		// workflow_owner           // offset 74, size 20
		// report_name              // offset 94, size  2
		assembly {
			// no shifting needed for bytes10 type
			workflowName := mload(add(metadata, 64))
			// shift right by 12 bytes to get the actual value
			workflowOwner := shr(mul(12, 8), mload(add(metadata, 74)))
		}
		return (workflowOwner, workflowName);
	}

	/// @notice Abstract function to process the report
	/// @param report The report calldata
	function _processReport(bytes calldata report) internal virtual;

	/// @inheritdoc IERC165
	function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
		return interfaceId == type(IReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
	}
}



// File: TemperatureConsumerAlpha.sol


pragma solidity ^0.8.26;


/// @notice Private Alpha-compatible consumer that skips metadata validation
contract AssetNavContract is IReceiverTemplate {
    // current NAV
    int256 public assetNAV;
    // Historic NAV
    struct NAVHistory {
        int256 nav;
        uint256 timestamp;
    }
    NAVHistory[] public navHistory;
    event AssetNavUpdated(int256 newAssetNav, uint256 timestamp);

    constructor(
        address expectedAuthor,
        bytes10 expectedWorkflowName
    ) IReceiverTemplate(expectedAuthor, expectedWorkflowName) {}

    function _processReport(bytes calldata report) internal override {
        int256 newAssetNav = abi.decode(report, (int256));
        assetNAV = newAssetNav;
        navHistory.push(NAVHistory({ nav: newAssetNav, timestamp: block.timestamp }));
        emit AssetNavUpdated(newAssetNav, block.timestamp);
    }

    function getNAVHistoryLength() external view returns (uint256) {
        return navHistory.length;
    }

    function getNAVAtIndex(
        uint256 index
    ) external view returns (int256 nav, uint256 timestamp) {
        require(index < navHistory.length, "Index out of bounds");
        NAVHistory storage entry = navHistory[index];
        return (entry.nav, entry.timestamp);
    }
    
    // Override to bypass metadata validation for MockKeystoneForwarder in Private Alpha
    function onReport(bytes calldata, bytes calldata report) external override {
        _processReport(report);
    }
}