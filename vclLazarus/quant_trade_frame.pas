unit quant_trade_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, PairSplitter, ExtCtrls, ComCtrls,
  StdCtrls;

type

  { TQuantTradeFrame }

  TQuantTradeFrame = class(TFrame)
    Edit1: TEdit;
    MainPanel: TPanel;
    MainChartPanel: TPanel;
    Panel1: TPanel;
    Splitter4: TSplitter;
    SubChartPanel: TPanel;
    PanelRight: TPanel;
    Panel4: TPanel;
    Panel3: TPanel;
    LeftPanel: TPanel;
    Splitter1: TSplitter;
    Splitter2: TSplitter;
    Splitter3: TSplitter;
    TopPanel: TPanel;
    Panel2: TPanel;
    BottomPanel: TPanel;
    procedure Edit1Change(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TQuantTradeFrame }

procedure TQuantTradeFrame.Edit1Change(Sender: TObject);
begin

end;

end.

