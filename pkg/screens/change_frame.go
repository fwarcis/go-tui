package screens

import "github.com/fwarcis/go-tui/pkg/frames"

func (s *Screen[WriterRespValue]) ChangeFrame(
	changeInTx func(frames.Mutable) error,
) error {
	s.frame.Lock()
	defer s.frame.Unlock()

	s.frame.WaitForWriting()

	err := changeInTx(&s.frame)
	if err != nil {
		return err
	}

	err = s.frame.Commit()
	if err != nil {
		return err
	}

	s.frame.SignalToWrite()

	return nil
}
