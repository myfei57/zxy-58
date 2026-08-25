package control

import (
	"errors"

	"poolops/internal/seq"
)

var errZoneRequired = errors.New("zone is required")

func (c *Controller) Treat(hour int, rawTurbidity float64) error {
	calibrated := c.Sensor.Read(rawTurbidity)
	_ = c.TurbFilter.Apply(calibrated)
	if !c.Schedule.Active(hour) {
		return nil
	}
	rec := seq.NewRecorder()
	c.runTreatment(rec)
	if err := c.Store.AppendHistory("treatment", rec.Last()); err != nil {
		return err
	}
	return nil
}

func (c *Controller) runTreatment(rec *seq.Recorder) {
	c.Speed.Increase()
	c.Circuit.Flow.SetRate(c.Speed.Current)
	c.PH.Stabilize()
	rec.Add("ph-stabilize")
	c.Doser.Apply()
	rec.Add("chlor-dose")
}

func (c *Controller) Backwash() error {
	rec := seq.NewRecorder()
	key := "backwash-" + c.FilterTank.ID
	err := c.retryBackwash(key, rec)
	if err != nil {
		return err
	}
	c.FilterTank.Rinse()
	c.Media.Age()
	c.FilterValve.SwitchToFilter()
	return c.Store.AppendHistory("backwash", rec.Last())
}

func (c *Controller) retryBackwash(key string, rec *seq.Recorder) error {
	return filterRetry(c.Store, key, func() error {
		return filterBackwash(c.FilterValve, c.DrainValve, rec)
	})
}

func (c *Controller) Fill(zoneID string) (float64, error) {
	if zoneID == "" {
		return 0, errZoneRequired
	}
	result, err := drainFill(c.Pool, zoneID, c.Hysteresis, c.Quota, c.Ledger)
	if err != nil {
		return 0, err
	}
	if result.Amount > 0 {
		c.Allocator.Consume(zoneID, uint64(result.Amount*1000))
	}
	return result.Amount, nil
}

func (c *Controller) Drain(zoneID string, level float64) error {
	if zoneID == "" {
		return errZoneRequired
	}
	rec := seq.NewRecorder()
	if err := drainRun(c.Store, c.Pool.ID, zoneID, level, rec); err != nil {
		return err
	}
	c.DrainValve.OpenDrain()
	return c.Store.AppendHistory("drain", rec.Last())
}

func (c *Controller) Failover() error {
	if err := pumpFailover(c.Roster, c.Circuit.Flow); err != nil {
		return err
	}
	c.PrimaryTel.RecordStop()
	c.StandbyTel.RecordStart()
	return nil
}

func (c *Controller) Restore(branches []string) error {
	rec := seq.NewRecorder()
	if err := circulateRestore(c.ReturnValves, branches, rec); err != nil {
		return err
	}
	return c.Store.AppendHistory("restore", rec.Last())
}

func (c *Controller) ResolveZone(probeID string) string {
	return c.Resolver.Resolve(c.Pool, probeID)
}
