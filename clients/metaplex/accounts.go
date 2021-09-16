// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AuctionManagerV1 struct {
	Key           Key
	Store         ag_solanago.PublicKey
	Authority     ag_solanago.PublicKey
	Auction       ag_solanago.PublicKey
	Vault         ag_solanago.PublicKey
	AcceptPayment ag_solanago.PublicKey
	State         AuctionManagerStateV1
	Settings      AuctionManagerSettingsV1

	// True if this is only winning configs of one item each, used for optimization in saving.
	StraightShotOptimization bool
}

func (obj AuctionManagerV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `Store` param:
	err = encoder.Encode(obj.Store)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Auction` param:
	err = encoder.Encode(obj.Auction)
	if err != nil {
		return err
	}
	// Serialize `Vault` param:
	err = encoder.Encode(obj.Vault)
	if err != nil {
		return err
	}
	// Serialize `AcceptPayment` param:
	err = encoder.Encode(obj.AcceptPayment)
	if err != nil {
		return err
	}
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	// Serialize `Settings` param:
	err = encoder.Encode(obj.Settings)
	if err != nil {
		return err
	}
	// Serialize `StraightShotOptimization` param:
	err = encoder.Encode(obj.StraightShotOptimization)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AuctionManagerV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `Store`:
	err = decoder.Decode(&obj.Store)
	if err != nil {
		return err
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Auction`:
	err = decoder.Decode(&obj.Auction)
	if err != nil {
		return err
	}
	// Deserialize `Vault`:
	err = decoder.Decode(&obj.Vault)
	if err != nil {
		return err
	}
	// Deserialize `AcceptPayment`:
	err = decoder.Decode(&obj.AcceptPayment)
	if err != nil {
		return err
	}
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	// Deserialize `Settings`:
	err = decoder.Decode(&obj.Settings)
	if err != nil {
		return err
	}
	// Deserialize `StraightShotOptimization`:
	err = decoder.Decode(&obj.StraightShotOptimization)
	if err != nil {
		return err
	}
	return nil
}