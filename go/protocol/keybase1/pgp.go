// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/pgp.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type SignMode int

const (
	SignMode_ATTACHED SignMode = 0
	SignMode_DETACHED SignMode = 1
	SignMode_CLEAR    SignMode = 2
)

func (o SignMode) DeepCopy() SignMode { return o }

var SignModeMap = map[string]SignMode{
	"ATTACHED": 0,
	"DETACHED": 1,
	"CLEAR":    2,
}

var SignModeRevMap = map[SignMode]string{
	0: "ATTACHED",
	1: "DETACHED",
	2: "CLEAR",
}

func (e SignMode) String() string {
	if v, ok := SignModeRevMap[e]; ok {
		return v
	}
	return ""
}

type PGPSignOptions struct {
	KeyQuery  string   `codec:"keyQuery" json:"keyQuery"`
	Mode      SignMode `codec:"mode" json:"mode"`
	BinaryIn  bool     `codec:"binaryIn" json:"binaryIn"`
	BinaryOut bool     `codec:"binaryOut" json:"binaryOut"`
}

func (o PGPSignOptions) DeepCopy() PGPSignOptions {
	return PGPSignOptions{
		KeyQuery:  o.KeyQuery,
		Mode:      o.Mode.DeepCopy(),
		BinaryIn:  o.BinaryIn,
		BinaryOut: o.BinaryOut,
	}
}

type PGPEncryptOptions struct {
	Recipients []string `codec:"recipients" json:"recipients"`
	NoSign     bool     `codec:"noSign" json:"noSign"`
	NoSelf     bool     `codec:"noSelf" json:"noSelf"`
	BinaryOut  bool     `codec:"binaryOut" json:"binaryOut"`
	KeyQuery   string   `codec:"keyQuery" json:"keyQuery"`
}

func (o PGPEncryptOptions) DeepCopy() PGPEncryptOptions {
	return PGPEncryptOptions{
		Recipients: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Recipients),
		NoSign:    o.NoSign,
		NoSelf:    o.NoSelf,
		BinaryOut: o.BinaryOut,
		KeyQuery:  o.KeyQuery,
	}
}

// PGPSigVerification is returned by pgpDecrypt and pgpVerify with information
// about a the signature verification. If isSigned is false, there was no
// signature, and the rest of the fields should be ignored.
type PGPSigVerification struct {
	IsSigned bool      `codec:"isSigned" json:"isSigned"`
	Verified bool      `codec:"verified" json:"verified"`
	Signer   User      `codec:"signer" json:"signer"`
	SignKey  PublicKey `codec:"signKey" json:"signKey"`
}

func (o PGPSigVerification) DeepCopy() PGPSigVerification {
	return PGPSigVerification{
		IsSigned: o.IsSigned,
		Verified: o.Verified,
		Signer:   o.Signer.DeepCopy(),
		SignKey:  o.SignKey.DeepCopy(),
	}
}

type PGPDecryptOptions struct {
	AssertSigned bool   `codec:"assertSigned" json:"assertSigned"`
	SignedBy     string `codec:"signedBy" json:"signedBy"`
}

func (o PGPDecryptOptions) DeepCopy() PGPDecryptOptions {
	return PGPDecryptOptions{
		AssertSigned: o.AssertSigned,
		SignedBy:     o.SignedBy,
	}
}

type PGPVerifyOptions struct {
	SignedBy  string `codec:"signedBy" json:"signedBy"`
	Signature []byte `codec:"signature" json:"signature"`
}

func (o PGPVerifyOptions) DeepCopy() PGPVerifyOptions {
	return PGPVerifyOptions{
		SignedBy: o.SignedBy,
		Signature: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Signature),
	}
}

type KeyInfo struct {
	Fingerprint string `codec:"fingerprint" json:"fingerprint"`
	Key         string `codec:"key" json:"key"`
	Desc        string `codec:"desc" json:"desc"`
}

func (o KeyInfo) DeepCopy() KeyInfo {
	return KeyInfo{
		Fingerprint: o.Fingerprint,
		Key:         o.Key,
		Desc:        o.Desc,
	}
}

type PGPQuery struct {
	Secret     bool   `codec:"secret" json:"secret"`
	Query      string `codec:"query" json:"query"`
	ExactMatch bool   `codec:"exactMatch" json:"exactMatch"`
}

func (o PGPQuery) DeepCopy() PGPQuery {
	return PGPQuery{
		Secret:     o.Secret,
		Query:      o.Query,
		ExactMatch: o.ExactMatch,
	}
}

type PGPCreateUids struct {
	UseDefault bool          `codec:"useDefault" json:"useDefault"`
	Ids        []PGPIdentity `codec:"ids" json:"ids"`
}

func (o PGPCreateUids) DeepCopy() PGPCreateUids {
	return PGPCreateUids{
		UseDefault: o.UseDefault,
		Ids: (func(x []PGPIdentity) []PGPIdentity {
			var ret []PGPIdentity
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Ids),
	}
}

// Export all pgp keys in lksec, then if doPurge is true, remove the keys from lksec.
type PGPPurgeRes struct {
	Filenames []string `codec:"filenames" json:"filenames"`
}

func (o PGPPurgeRes) DeepCopy() PGPPurgeRes {
	return PGPPurgeRes{
		Filenames: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Filenames),
	}
}

type PGPSignArg struct {
	SessionID int            `codec:"sessionID" json:"sessionID"`
	Source    Stream         `codec:"source" json:"source"`
	Sink      Stream         `codec:"sink" json:"sink"`
	Opts      PGPSignOptions `codec:"opts" json:"opts"`
}

func (o PGPSignArg) DeepCopy() PGPSignArg {
	return PGPSignArg{
		SessionID: o.SessionID,
		Source:    o.Source.DeepCopy(),
		Sink:      o.Sink.DeepCopy(),
		Opts:      o.Opts.DeepCopy(),
	}
}

type PGPPullArg struct {
	SessionID   int      `codec:"sessionID" json:"sessionID"`
	UserAsserts []string `codec:"userAsserts" json:"userAsserts"`
}

func (o PGPPullArg) DeepCopy() PGPPullArg {
	return PGPPullArg{
		SessionID: o.SessionID,
		UserAsserts: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.UserAsserts),
	}
}

type PGPEncryptArg struct {
	SessionID int               `codec:"sessionID" json:"sessionID"`
	Source    Stream            `codec:"source" json:"source"`
	Sink      Stream            `codec:"sink" json:"sink"`
	Opts      PGPEncryptOptions `codec:"opts" json:"opts"`
}

func (o PGPEncryptArg) DeepCopy() PGPEncryptArg {
	return PGPEncryptArg{
		SessionID: o.SessionID,
		Source:    o.Source.DeepCopy(),
		Sink:      o.Sink.DeepCopy(),
		Opts:      o.Opts.DeepCopy(),
	}
}

type PGPDecryptArg struct {
	SessionID int               `codec:"sessionID" json:"sessionID"`
	Source    Stream            `codec:"source" json:"source"`
	Sink      Stream            `codec:"sink" json:"sink"`
	Opts      PGPDecryptOptions `codec:"opts" json:"opts"`
}

func (o PGPDecryptArg) DeepCopy() PGPDecryptArg {
	return PGPDecryptArg{
		SessionID: o.SessionID,
		Source:    o.Source.DeepCopy(),
		Sink:      o.Sink.DeepCopy(),
		Opts:      o.Opts.DeepCopy(),
	}
}

type PGPVerifyArg struct {
	SessionID int              `codec:"sessionID" json:"sessionID"`
	Source    Stream           `codec:"source" json:"source"`
	Opts      PGPVerifyOptions `codec:"opts" json:"opts"`
}

func (o PGPVerifyArg) DeepCopy() PGPVerifyArg {
	return PGPVerifyArg{
		SessionID: o.SessionID,
		Source:    o.Source.DeepCopy(),
		Opts:      o.Opts.DeepCopy(),
	}
}

type PGPImportArg struct {
	SessionID  int    `codec:"sessionID" json:"sessionID"`
	Key        []byte `codec:"key" json:"key"`
	PushSecret bool   `codec:"pushSecret" json:"pushSecret"`
}

func (o PGPImportArg) DeepCopy() PGPImportArg {
	return PGPImportArg{
		SessionID: o.SessionID,
		Key: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Key),
		PushSecret: o.PushSecret,
	}
}

type PGPExportArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Options   PGPQuery `codec:"options" json:"options"`
}

func (o PGPExportArg) DeepCopy() PGPExportArg {
	return PGPExportArg{
		SessionID: o.SessionID,
		Options:   o.Options.DeepCopy(),
	}
}

type PGPExportByFingerprintArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Options   PGPQuery `codec:"options" json:"options"`
}

func (o PGPExportByFingerprintArg) DeepCopy() PGPExportByFingerprintArg {
	return PGPExportByFingerprintArg{
		SessionID: o.SessionID,
		Options:   o.Options.DeepCopy(),
	}
}

type PGPExportByKIDArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Options   PGPQuery `codec:"options" json:"options"`
}

func (o PGPExportByKIDArg) DeepCopy() PGPExportByKIDArg {
	return PGPExportByKIDArg{
		SessionID: o.SessionID,
		Options:   o.Options.DeepCopy(),
	}
}

type PGPKeyGenArg struct {
	SessionID   int           `codec:"sessionID" json:"sessionID"`
	PrimaryBits int           `codec:"primaryBits" json:"primaryBits"`
	SubkeyBits  int           `codec:"subkeyBits" json:"subkeyBits"`
	CreateUids  PGPCreateUids `codec:"createUids" json:"createUids"`
	AllowMulti  bool          `codec:"allowMulti" json:"allowMulti"`
	DoExport    bool          `codec:"doExport" json:"doExport"`
	PushSecret  bool          `codec:"pushSecret" json:"pushSecret"`
}

func (o PGPKeyGenArg) DeepCopy() PGPKeyGenArg {
	return PGPKeyGenArg{
		SessionID:   o.SessionID,
		PrimaryBits: o.PrimaryBits,
		SubkeyBits:  o.SubkeyBits,
		CreateUids:  o.CreateUids.DeepCopy(),
		AllowMulti:  o.AllowMulti,
		DoExport:    o.DoExport,
		PushSecret:  o.PushSecret,
	}
}

type PGPKeyGenDefaultArg struct {
	SessionID  int           `codec:"sessionID" json:"sessionID"`
	CreateUids PGPCreateUids `codec:"createUids" json:"createUids"`
}

func (o PGPKeyGenDefaultArg) DeepCopy() PGPKeyGenDefaultArg {
	return PGPKeyGenDefaultArg{
		SessionID:  o.SessionID,
		CreateUids: o.CreateUids.DeepCopy(),
	}
}

type PGPDeletePrimaryArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o PGPDeletePrimaryArg) DeepCopy() PGPDeletePrimaryArg {
	return PGPDeletePrimaryArg{
		SessionID: o.SessionID,
	}
}

type PGPSelectArg struct {
	SessionID        int    `codec:"sessionID" json:"sessionID"`
	FingerprintQuery string `codec:"fingerprintQuery" json:"fingerprintQuery"`
	AllowMulti       bool   `codec:"allowMulti" json:"allowMulti"`
	SkipImport       bool   `codec:"skipImport" json:"skipImport"`
	OnlyImport       bool   `codec:"onlyImport" json:"onlyImport"`
}

func (o PGPSelectArg) DeepCopy() PGPSelectArg {
	return PGPSelectArg{
		SessionID:        o.SessionID,
		FingerprintQuery: o.FingerprintQuery,
		AllowMulti:       o.AllowMulti,
		SkipImport:       o.SkipImport,
		OnlyImport:       o.OnlyImport,
	}
}

type PGPUpdateArg struct {
	SessionID    int      `codec:"sessionID" json:"sessionID"`
	All          bool     `codec:"all" json:"all"`
	Fingerprints []string `codec:"fingerprints" json:"fingerprints"`
}

func (o PGPUpdateArg) DeepCopy() PGPUpdateArg {
	return PGPUpdateArg{
		SessionID: o.SessionID,
		All:       o.All,
		Fingerprints: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Fingerprints),
	}
}

type PGPPurgeArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	DoPurge   bool `codec:"doPurge" json:"doPurge"`
}

func (o PGPPurgeArg) DeepCopy() PGPPurgeArg {
	return PGPPurgeArg{
		SessionID: o.SessionID,
		DoPurge:   o.DoPurge,
	}
}

type PGPStorageDismissArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

func (o PGPStorageDismissArg) DeepCopy() PGPStorageDismissArg {
	return PGPStorageDismissArg{
		SessionID: o.SessionID,
	}
}

type PGPInterface interface {
	PGPSign(context.Context, PGPSignArg) error
	// Download PGP keys for tracked users and update the local GPG keyring.
	// If usernames is nonempty, update only those users.
	PGPPull(context.Context, PGPPullArg) error
	PGPEncrypt(context.Context, PGPEncryptArg) error
	PGPDecrypt(context.Context, PGPDecryptArg) (PGPSigVerification, error)
	PGPVerify(context.Context, PGPVerifyArg) (PGPSigVerification, error)
	PGPImport(context.Context, PGPImportArg) error
	// Exports active PGP keys. Only allows armored export.
	PGPExport(context.Context, PGPExportArg) ([]KeyInfo, error)
	PGPExportByFingerprint(context.Context, PGPExportByFingerprintArg) ([]KeyInfo, error)
	PGPExportByKID(context.Context, PGPExportByKIDArg) ([]KeyInfo, error)
	PGPKeyGen(context.Context, PGPKeyGenArg) error
	PGPKeyGenDefault(context.Context, PGPKeyGenDefaultArg) error
	PGPDeletePrimary(context.Context, int) error
	// Select an existing key and add to Keybase.
	PGPSelect(context.Context, PGPSelectArg) error
	// Push updated key(s) to the server.
	PGPUpdate(context.Context, PGPUpdateArg) error
	PGPPurge(context.Context, PGPPurgeArg) (PGPPurgeRes, error)
	// Dismiss the PGP unlock via secret_store_file notification.
	PGPStorageDismiss(context.Context, int) error
}

func PGPProtocol(i PGPInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.pgp",
		Methods: map[string]rpc.ServeHandlerDescription{
			"pgpSign": {
				MakeArg: func() interface{} {
					ret := make([]PGPSignArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPSignArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPSignArg)(nil), args)
						return
					}
					err = i.PGPSign(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpPull": {
				MakeArg: func() interface{} {
					ret := make([]PGPPullArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPPullArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPPullArg)(nil), args)
						return
					}
					err = i.PGPPull(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpEncrypt": {
				MakeArg: func() interface{} {
					ret := make([]PGPEncryptArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPEncryptArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPEncryptArg)(nil), args)
						return
					}
					err = i.PGPEncrypt(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpDecrypt": {
				MakeArg: func() interface{} {
					ret := make([]PGPDecryptArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPDecryptArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPDecryptArg)(nil), args)
						return
					}
					ret, err = i.PGPDecrypt(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpVerify": {
				MakeArg: func() interface{} {
					ret := make([]PGPVerifyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPVerifyArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPVerifyArg)(nil), args)
						return
					}
					ret, err = i.PGPVerify(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpImport": {
				MakeArg: func() interface{} {
					ret := make([]PGPImportArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPImportArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPImportArg)(nil), args)
						return
					}
					err = i.PGPImport(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpExport": {
				MakeArg: func() interface{} {
					ret := make([]PGPExportArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPExportArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPExportArg)(nil), args)
						return
					}
					ret, err = i.PGPExport(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpExportByFingerprint": {
				MakeArg: func() interface{} {
					ret := make([]PGPExportByFingerprintArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPExportByFingerprintArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPExportByFingerprintArg)(nil), args)
						return
					}
					ret, err = i.PGPExportByFingerprint(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpExportByKID": {
				MakeArg: func() interface{} {
					ret := make([]PGPExportByKIDArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPExportByKIDArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPExportByKIDArg)(nil), args)
						return
					}
					ret, err = i.PGPExportByKID(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpKeyGen": {
				MakeArg: func() interface{} {
					ret := make([]PGPKeyGenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPKeyGenArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPKeyGenArg)(nil), args)
						return
					}
					err = i.PGPKeyGen(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpKeyGenDefault": {
				MakeArg: func() interface{} {
					ret := make([]PGPKeyGenDefaultArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPKeyGenDefaultArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPKeyGenDefaultArg)(nil), args)
						return
					}
					err = i.PGPKeyGenDefault(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpDeletePrimary": {
				MakeArg: func() interface{} {
					ret := make([]PGPDeletePrimaryArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPDeletePrimaryArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPDeletePrimaryArg)(nil), args)
						return
					}
					err = i.PGPDeletePrimary(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpSelect": {
				MakeArg: func() interface{} {
					ret := make([]PGPSelectArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPSelectArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPSelectArg)(nil), args)
						return
					}
					err = i.PGPSelect(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpUpdate": {
				MakeArg: func() interface{} {
					ret := make([]PGPUpdateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPUpdateArg)(nil), args)
						return
					}
					err = i.PGPUpdate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpPurge": {
				MakeArg: func() interface{} {
					ret := make([]PGPPurgeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPPurgeArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPPurgeArg)(nil), args)
						return
					}
					ret, err = i.PGPPurge(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pgpStorageDismiss": {
				MakeArg: func() interface{} {
					ret := make([]PGPStorageDismissArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PGPStorageDismissArg)
					if !ok {
						err = rpc.NewTypeError((*[]PGPStorageDismissArg)(nil), args)
						return
					}
					err = i.PGPStorageDismiss(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type PGPClient struct {
	Cli rpc.GenericClient
}

func (c PGPClient) PGPSign(ctx context.Context, __arg PGPSignArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpSign", []interface{}{__arg}, nil)
	return
}

// Download PGP keys for tracked users and update the local GPG keyring.
// If usernames is nonempty, update only those users.
func (c PGPClient) PGPPull(ctx context.Context, __arg PGPPullArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpPull", []interface{}{__arg}, nil)
	return
}

func (c PGPClient) PGPEncrypt(ctx context.Context, __arg PGPEncryptArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpEncrypt", []interface{}{__arg}, nil)
	return
}

func (c PGPClient) PGPDecrypt(ctx context.Context, __arg PGPDecryptArg) (res PGPSigVerification, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpDecrypt", []interface{}{__arg}, &res)
	return
}

func (c PGPClient) PGPVerify(ctx context.Context, __arg PGPVerifyArg) (res PGPSigVerification, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpVerify", []interface{}{__arg}, &res)
	return
}

func (c PGPClient) PGPImport(ctx context.Context, __arg PGPImportArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpImport", []interface{}{__arg}, nil)
	return
}

// Exports active PGP keys. Only allows armored export.
func (c PGPClient) PGPExport(ctx context.Context, __arg PGPExportArg) (res []KeyInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpExport", []interface{}{__arg}, &res)
	return
}

func (c PGPClient) PGPExportByFingerprint(ctx context.Context, __arg PGPExportByFingerprintArg) (res []KeyInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpExportByFingerprint", []interface{}{__arg}, &res)
	return
}

func (c PGPClient) PGPExportByKID(ctx context.Context, __arg PGPExportByKIDArg) (res []KeyInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpExportByKID", []interface{}{__arg}, &res)
	return
}

func (c PGPClient) PGPKeyGen(ctx context.Context, __arg PGPKeyGenArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpKeyGen", []interface{}{__arg}, nil)
	return
}

func (c PGPClient) PGPKeyGenDefault(ctx context.Context, __arg PGPKeyGenDefaultArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpKeyGenDefault", []interface{}{__arg}, nil)
	return
}

func (c PGPClient) PGPDeletePrimary(ctx context.Context, sessionID int) (err error) {
	__arg := PGPDeletePrimaryArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpDeletePrimary", []interface{}{__arg}, nil)
	return
}

// Select an existing key and add to Keybase.
func (c PGPClient) PGPSelect(ctx context.Context, __arg PGPSelectArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpSelect", []interface{}{__arg}, nil)
	return
}

// Push updated key(s) to the server.
func (c PGPClient) PGPUpdate(ctx context.Context, __arg PGPUpdateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpUpdate", []interface{}{__arg}, nil)
	return
}

func (c PGPClient) PGPPurge(ctx context.Context, __arg PGPPurgeArg) (res PGPPurgeRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpPurge", []interface{}{__arg}, &res)
	return
}

// Dismiss the PGP unlock via secret_store_file notification.
func (c PGPClient) PGPStorageDismiss(ctx context.Context, sessionID int) (err error) {
	__arg := PGPStorageDismissArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.pgp.pgpStorageDismiss", []interface{}{__arg}, nil)
	return
}
