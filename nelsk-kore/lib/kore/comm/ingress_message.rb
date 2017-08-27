module Kore
  module Comm
    class IngressMessage
      attr_accessor :originator, :plugin, :cmd, :content, :raw
      def initialize(o)
        @originator = o.fetch(:originator, '')
        @plugin = o.fetch(:plugin, '')
        @cmd = o.fetch(:cmd, '')
        @content = o.fetch(:content, '')
        @raw = o.fetch(:raw, '')
      end

      def to_s
        "[#{@originator}] [#{@plugin}] [#{@cmd}] [#{@content}]"
      end
    end
  end
end
